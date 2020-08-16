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
                sh 'curl -s -o --request GET curl --location --request GET 52.201.125.229/call/teste6/sendMail?emails=rrms@cin.ufpe.br'
            }
        }
    }
}
