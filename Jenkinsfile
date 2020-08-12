pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                echo 'Building..'
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
                sh 'curl -s -o --request GET 127.0.0.1/call/enviaremail10/sendMail?emails=rcrdrobson@gmail.com,rrms@cin.ufpe.br'
            }
        }
    }
}