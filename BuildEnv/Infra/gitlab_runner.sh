curl -L https://packages.gitlab.com/install/repositories/runner/gitlab-runner/script.deb.sh | bash && apt-get update && apt-get install gitlab-runner
host_name=`hostname`
gitlab-runner register -n  -u https://gitlab.com/ -r $gitlab_token --executor shell --tag-list linux
echo "gitlab-runner ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/gitlab
gitlab-runner start
