pipeline {
  agent any
  stages {
    stage('gitcode') {
      steps {
        git(url: 'https://github.com/frankwsj/go_code.git', branch: 'master', credentialsId: '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', poll: true)
      }
    }

  }
}