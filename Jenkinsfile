pipeline {
    agent { docker { image 'python:3.7.1' } }
    stages {
        stage('Build') {
            steps {
            	sh 'go build -o essai-service main.go'
            }
        }
    }
}
