curl -L https://packages.gitlab.com/install/repositories/runner/gitlab-runner/script.deb.sh | bash && apt-get update && apt-get install gitlab-runner
host_name=`hostname`
gitlab-runner register --non-interactive --locked false --run-untagged true --tag-list linux --name $host_name --registration-token {{ gitlab_token }} --url https://gitlab.com/ --executor shell
echo "gitlab-runner ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/gitlab
gitlab-runner start