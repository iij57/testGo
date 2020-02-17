pipeline {
  agent {
    dockerfile {
      filename 'Dockerfile'
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