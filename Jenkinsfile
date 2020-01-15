pipeline {
  agent {
    docker {
      image '1.12-stretch'
    }

  }
  stages {
    stage('build') {
      steps {
        sh 'go build'
      }
    }

  }
}