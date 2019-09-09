pipeline {
    agent none
    environment {
        UPANGU_AXE = 'dev'
    }
    stages {
        stage('Build') {
            agent {
                docker {
                    image 'golang:1.12-alpine'
                }
            }
            steps {
                sh 'sh ./scripts/build.sh'
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
                sh 'sh ./scripts/deploy.sh development'
                input message: ''
                sh 'sh ./scripts/manage.sh stop'
            }
        }
        stage('Deploy for production') {
            when {
                branch 'production'
            }
            steps {
                sh 'sh ./scripts/deploy.sh production'
                input message: ''
                sh 'sh ./scripts/manage.sh stop'
            }
        }
    }
}
