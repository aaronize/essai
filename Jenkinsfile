Jenkinsfile (Declarative Pipeline)
pipeline {
    agent { docker { image 'python:3.7.1' } }
    stages {
        stage('Build') {
            steps {
                retry (3) {

                }
                sh 'flask run'
            }
        }
    }
}