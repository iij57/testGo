pipeline {
  agent {
    docker {
      image 'golang'
    }

  }
  stages {
    stage('Build') {
      steps {
        sh 'go get -u github.com/julienschmidt/httprouter'
        sh 'go build .'
      }
    }

  }
}