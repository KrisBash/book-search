import React, { Component } from 'react';
import logo from '../books-banner.jpg';
import axios from 'axios';
import './App.css';

export const BookDetail = (props) => {
  //console.log(props);
  return (
     <div> {props.book.title} </div>
  );
};


export class BookForm extends React.Component {
  constructor() {
      super();
      this.state = {
          books: [],
          book_title:"",
          book_authors:"",
          book_isbn:"",
          book_publisher:"",
          book_published_date:"",
          book_page_count:"",
          book_description:"",
          book_print_type:"",
          search_isbn: "",
          cb_book_title:true,
          cb_book_authors:true,
          cb_book_publisher:false,
          cb_book_published_date:false,
          cb_book_page_count:true,
          cb_book_description:true,          
          cb_book_average_rating:false,        
          cb_book_print_type:false

      };
      this.handleChange = this.handleChange.bind(this);
      this.handleSubmit = this.handleSubmit.bind(this);      
  }

  handleChange(event) {
    //console.log(event.target.name + " - " + event.target.value)
    const target = event.target;
    const value = target.type === 'checkbox' ? target.checked : target.value;
    const name = target.name;

    this.setState({
      [name]: value
    });
  }

  search_for_book(isbn) {
    const api_host = process.env.BOOK_API_SERVICE_HOST
    const api_port = process.env.BOOK_API_SERVICE_PORT
    console.log("api_host = " + api_host)
    console.log("api_port = " + api_port)
    this.url_root = `http://10.0.229.194:8222/graphql?query={book(isbn:"` + isbn + `"){isbn,`;
    this.fields = "";
    this.url_tail = `}}`;
    if (this.state.cb_book_authors){this.fields += "authors,"};
    if (this.state.cb_book_title){this.fields += "title,"};
    if (this.state.cb_book_publisher){this.fields += "publisher,"};
    if (this.state.cb_book_published_date){this.fields += "published_date,"};
    if (this.state.cb_book_page_count){this.fields += "page_count,"};
    if (this.state.cb_book_description){this.fields += "description,"};
    if (this.state.cb_book_average_rating){this.fields += "average_rating,"};
    if (this.state.cb_book_print_type){this.fields += "print_type,"};
    this.fields  =  this.fields.substring(0, this.fields.length - 1);
    this.apiUrl = this.url_root + this.fields + this.url_tail

    return  axios.get(this.apiUrl,{
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
        'Accept': 'text/json'
      }}).then((response) => {
        //console.log(response);
          var book_isbn = response.data.data.book.isbn;
          var book_title = response.data.data.book.title;
          var book_authors = response.data.data.book.authors;
          var book_publisher = response.data.data.book.publisher;
          var book_published_date = response.data.data.book.published_date;
          var book_page_count  = response.data.data.book.page_count;
          var book_description = response.data.data.book.description;
          var book_print_type = response.data.data.book.print_type;
          this.setState({book_title: book_title});
          this.setState({book_isbn: book_isbn});
          this.setState({book_authors: book_authors});
          this.setState({book_publisher: book_publisher});
          this.setState({book_published_date: book_published_date});
          this.setState({book_page_count: book_page_count});
          this.setState({book_description: book_description});
          this.setState({book_publisher: book_publisher});
          this.setState({book_print_type: book_print_type});
          //console.log(this.state.book_title)  
      })
      .catch(error => { console.log('error', error); });
  }
  
  showBooks() {
      //console.log(this.state.books.map);
      return this.state.books.map(book => <BookDetail book={book}  />);
      }
  
  handleSubmit(event) {
  
    if (Number(this.state.search_isbn)> 0){
      this.search_for_book(this.state.search_isbn);
    }else{
      alert("Please enter a numeric ISBN")
    }
    event.preventDefault();
  }

  render() {
      return (
          <div>
            <div className="search_form">
              <form onSubmit={this.handleSubmit} id="cl_search_form">
                <label>
                  ISBN:
                  <input type="string" value={this.state.search_isbn} name="search_isbn" onChange={this.handleChange} />
                </label>
                <input type="submit" value="Submit" />
                <br/>
                <blockquote>
                <b>Display:</b>
                <label>
                  Title:
                  <input name="cb_book_title" type="checkbox" checked={this.state.cb_book_title} onChange={this.handleChange} />
                </label>
                &nbsp;
                <label>
                  Author:
                  <input name="cb_book_authors" type="checkbox" checked={this.state.cb_book_authors} onChange={this.handleChange} /> 
                </label>
                &nbsp;
                <label>
                  Publisher:
                  <input name="cb_book_publisher" type="checkbox" checked={this.state.cb_book_publisher} onChange={this.handleChange} />
                </label>
                &nbsp;
                <label>
                  Published Date:
                  <input name="cb_book_published_date" type="checkbox" checked={this.state.cb_book_published_date} onChange={this.handleChange} />
                </label>   
                &nbsp;       
                <label>
                  Page Count:
                  <input name="cb_book_page_count" type="checkbox" checked={this.state.cb_book_page_count} onChange={this.handleChange} />
                </label>
                &nbsp;      
                <label>
                  Description:
                  <input name="cb_book_description" type="checkbox" checked={this.state.cb_book_description} onChange={this.handleChange} />
                </label>   
                &nbsp;
                <label>
                  Print Type:
                  <input
                    name="cb_book_print_type" type="checkbox"
                    checked={this.state.cb_book_print_type}
                    onChange={this.handleChange} />
                </label>                             
                </blockquote>                
              </form>
            </div>
            <div>
              <table className="result_table">
                <tbody>
                <tr><td width="100px"><b>ISBN:</b></td><td>{this.state.book_isbn}</td></tr>
                <tr><td><b>Title:</b></td><td>{this.state.book_title}</td></tr>
                <tr><td><b>Authors:</b></td><td>{this.state.book_authors}</td></tr>
                <tr><td><b>Publisher:</b></td><td>{this.state.book_publisher}</td></tr>
                <tr><td><b>Published Date:</b></td><td>{this.state.book_published_date}</td></tr>
                <tr><td><b>Page Count:</b></td><td>{this.state.book_page_count}</td></tr>
                <tr><td><b>Print Type:</b></td><td>{this.state.book_print_type}</td></tr>
                <tr><td colSpan="2"><b>Description</b></td></tr>
                <tr><td colSpan="2">{this.state.book_description}</td></tr>
                </tbody>
              </table>
            </div>
          </div>
      );
  }

}


class Header extends Component {
  render() {
    return (
      <div className="App">
        <div className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
        </div>
        <div className="page_header">
        <div className="App-intro">
            <b>Search for a book by ISBN</b>
            <div className="body_text">
              <i>Some example ISBNs: 1491935677 or 9780545010221</i>
            </div>
        </div>
        </div>
      </div>

    );
  }
}

class App extends Component {
  componentWillMount(){
    const api_host = process.env.BOOK_API_SERVICE_HOST
    const api_port = process.env.BOOK_API_SERVICE_PORT
    console.log ("host" + api_host)
    console.log ("port" + api_port)
  }
  componentDidMount() {
    
  }
  render() {
    return (
      
      <div>
        <Header/>
        <BookForm/>
      </div>
    );
  }
}


export default App;
