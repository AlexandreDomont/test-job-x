pipeline {
    agent any
    stages {
        stage('NORMAL Stage') {
            steps {
                echo 'I am one'
            }
        }
        stage('Parallel Stage') {
           parallel {
                stage('stage one') {
                    steps {
                        echo "Me in stage one"
                    }
                }
                stage('Stage two') {
                    steps {
                        echo "Me in stage two"
                    }
                }
                }
            }
        }
}
