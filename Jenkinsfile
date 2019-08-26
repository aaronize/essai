pipeline {
    agent { docker { image 'python:3.7.1' } }
    stages {
        stage('Build') {
            docker {
		image "golang:latest"
	    }
            steps {
            	sh 'go build -o essai-api main.go'
            }
        }
    }
}
