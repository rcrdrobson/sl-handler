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
                curl --location --request POST 'localhost/call/enviaremail10/sendMail?emails=rcrdrobson@gmail.com,rrms@cin.ufpe.br'
            }
        }
    }
}