pipeline {
  agent any
  stages {
    stage('PreTest') {
      steps {
        bat 'go get -u github.com/julienschmidt/httprouter'
      }
    }

    stage('Build') {
      steps {
        bat 'cd src/main && go build .'
      }
    }

    stage('Test') {
      steps {
        bat 'cd src/test && go test'
      }
    }

  }
}