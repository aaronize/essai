pipeline {
    agent {
        docker {
            image 'essai-api:latest'
            args '-d --rm --name upangu-axe -p 15101:15101 -p 15102:15102'
        }
    }
    environment {
        UPANGU_AXE = 'dev'
    }
    stages {
        stage('Build') {
            steps {
                sh './scripts/build.sh'
            }
        }
        stage('Test') {
            steps {
                sh 'echo "testing"'
            }
        }
        stage('Delivery for development') {
            when {
                branch 'develop'
            }
            steps {
                sh './scripts/deploy.sh development'
                input message: ''
                sh './scripts/manage.sh stop'
            }
        }
        stage('Deploy for production') {
            when {
                branch 'production'
            }
            steps {
                sh './scripts/deploy.sh production'
                input message: ''
                sh './scripts/manage.sh stop'
            }
        }
    }
}
