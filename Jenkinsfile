pipeline {
  agent {
    docker {
      image '1.12.15-stretch'
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