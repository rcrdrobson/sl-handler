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
                sh 'curl --location --request POST http://177.25.86.35/call/enviaremail10/sendMail?emails=rcrdrobson@gmail.com,rrms@cin.ufpe.br'
            }
        }
    }
}