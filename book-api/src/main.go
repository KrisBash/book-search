package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"net/http"
	"time"
	"os"
	"github.com/graphql-go/graphql"
	"github.com/go-redis/redis" 	
	log "github.com/sirupsen/logrus"	
//	consulapi "github.com/hashicorp/consul/api"
)

var redis_cl  = redis.NewClient(&redis.Options{
	Addr:     cachedb_address,
	Password: "", // no password set
	DB:       0,  // use default DB
})

type book_sparse struct {
	isbn       string `json:"isbn"`
	title      string `json:"title"`
	authors    string `json:"authors"`
	description     string `json:"description"`
	published_date string `json:"publisehd_date"`
	publisher       string  `json:"publisher"`
	print_type      string  `"print_type"`
	average_rating float64  `"average_rating"`
	image_links     string  `"image_links"`
	page_count      int             `page_count`
}

var book_type = graphql.NewObject(
	graphql.ObjectConfig{
			Name: "Book",
			Fields: graphql.Fields{
					"isbn": &graphql.Field{
							Type: graphql.String,
					},
					"title": &graphql.Field{
							Type: graphql.String,
					},
					"authors": &graphql.Field{
							Type: graphql.String,
					},
					"description": &graphql.Field{
							Type: graphql.String,
					},
					"page_count": &graphql.Field{Type: graphql.Int,},
					"published_date": &graphql.Field{Type: graphql.String,},
					"publisher": &graphql.Field{Type: graphql.String,},
					"print_type": &graphql.Field{Type: graphql.String,},
					"average_rating": &graphql.Field{Type: graphql.String,},
					"image_links": &graphql.Field{Type: graphql.String,},
			},
	},
)

func cache_set(key string, value string){
	log.Info("Caching record: ", value)
	err := redis_cl.Set(key, value, time.Hour *12).Err()
	if err != nil {
		log.Error("Failed to set cache.",err)
	}

}


var queryType = graphql.NewObject(
	graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
					"book": &graphql.Field{
							Type: book_type,
							Args: graphql.FieldConfigArgument{
									"isbn": &graphql.ArgumentConfig{
											Type: graphql.String,
									},
							},

							Resolve: func(p graphql.ResolveParams) (interface{}, error) {
									search_isbn := p.Args["isbn"].(string)
									statsd_incr("book_requests")
									log.Info("Searching for a book with isbn: ", search_isbn)

									if search_isbn == "test"{
											log.Info("isbn is set to test")
											volume :=  map[string]interface{}{
													"isbn":       fmt.Sprintf("test"),
													"authors":     fmt.Sprintf("test"),
													"title":          fmt.Sprintf("test"),
													"description":    fmt.Sprintf("test"),
													"published_date":        fmt.Sprintf("test"),
													"publisher":      fmt.Sprintf("test"),
													"print_type":     fmt.Sprintf("test"),
													"average_rating":       fmt.Sprintf("%v\n",1),
													"image_links":          fmt.Sprintf("test"),
													"page_count":  fmt.Sprintf("%v\n",1),
											}
											return volume, nil
									}

									//query cache first
									val, err := redis_cl.Get(search_isbn).Result()
									if err != nil {
										statsd_incr("cache_miss")
										log.Info("No result in cache for ", search_isbn)
									}else{
										statsd_incr("cache_hit")
										log.Info("Found a cache record for: ",search_isbn)
										var result map[string]interface{}
										json.Unmarshal([]byte(val), &result)
																			
										volume :=  map[string]interface{}{
											"isbn":       result["isbn"],
											"authors":     result["authors"],
											"title":          result["title"],
											"description":    result["description"],
											"published_date":        result["published_date"],
											"publisher":      result["publisher"],
											"print_type":     result["print_type"],
											"average_rating":       result["average_rating"],
											"image_links":          result["image_links"],
											"page_count":  result["page_count"],
										}
										return volume, nil										
											
									}
									

									var book_record book_volume
									start := time.Now()
									book_record = get_book_by_isbn(search_isbn)

									//measure and emit lookup time
									elapsed := time.Since(start).Seconds()
									statsd_gauge("book_search_time_s",elapsed)
									
									
									items := book_record.Items[0]
									volume_info := items.VolumeInfo
									image_links := volume_info.ImageLinks

									log.Info("Book found:  ", book_record.TotalItems)
									if book_record.TotalItems < 1 {
											log.Warn("No records found for isbn")
											return nil, nil
									} else {
											log.Info("Book found:  ", book_record.TotalItems)

											volume :=  map[string]interface{}{
													"isbn":       fmt.Sprintf("%v",search_isbn),
													"authors":     fmt.Sprintf(strings.Join(volume_info.Authors, ", ")),
													"title":          fmt.Sprintf(volume_info.Title),
													"description":    fmt.Sprintf(volume_info.Description),
													"published_date":         fmt.Sprintf(volume_info.PublishedDate),
													"publisher":      fmt.Sprintf(volume_info.Publisher),
													"print_type":     fmt.Sprintf(volume_info.PrintType),
													"average_rating":       fmt.Sprintf("%v",volume_info.AverageRating),
													"image_links":          fmt.Sprintf(image_links.SmallThumbnail),
													"page_count":  fmt.Sprintf("%v",volume_info.PageCount),
											}
											
											//marshal book info to JSON for cache
											book_json,err := json.Marshal(volume)
											if err != nil {
												log.Error("Unable to marshal book record to JSON", err)
											}
											cache_set(search_isbn, string(book_json))
											return volume, nil
											}

							},
					},
			},
	},
)


var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
			Query: queryType,
	},
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	var cachedb_server string =  os.Getenv("CACHE_DB_SERVICE_HOST")
	var cachedb_port string =  os.Getenv("CACHE_DB_SERVICE_PORT")
	cachedb_address := cachedb_server + ":" + cachedb_port

	pong, err := redis_cl.Ping().Result()
	if err != nil {
			log.Error(err)
	}else{
			//statsd_gauge("redis_ping",pong)
			log.Info(pong)
	}

	// Output: PONG <nil>

/*      var consul_server string =  os.Getenv("CONSUL_SERVER")
	var consul_port string =  os.Getenv("CONSUL_PORT")
	config := consulapi.DefaultConfig()
	consul_address := consul_server + ":" + consul_port
	config.Address = "192.168.1.2:8500"
	consul, err := consulapi.NewClient(config)
*/
}


func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: query,
	})
	if len(result.Errors) > 0 {
			log.Error("wrong result, unexpected errors: %v. Isbn may be wrong", result.Errors)
	}
	return result
}

func main() {
	//_ = importJSONDataFromFile("data.json", &data)
	//test

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		
			result := executeQuery(r.URL.Query().Get("query"), schema)
			json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Now server is running on port 8222")
	//fmt.Println("Test with Get      : curl -g 'http://localhost:8080/graphql?query={user(id:\"1\"){name}}'")
	http.ListenAndServe(":8222", nil)
}
