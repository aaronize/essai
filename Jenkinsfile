pipeline {
    agent any
    environment {
        UPANGU_AXE = 'dev'
    }
    stages {
        stage('GoBuild') {
            agent {
                docker {
                    image 'golang:latest'
                    args '-v /data/jenkins_home/workspace/essai-api:/go/src/essai'
                }
            }
            steps {
                sh 'bash ./scripts/build.sh golang'
            }
        }
        stage('DockerBuild') {
            steps {
                sh 'sh ./scripts/build.sh docker'
            }
        }
        stage('Test') {
            steps {
                sh 'echo "testing"'
            }
        }
        stage('Deploy') {
            steps {
                sh 'sh ./scripts/deploy.sh development'
            }
        }
        stage('Deploy for development') {
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
