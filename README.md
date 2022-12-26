# Understanding Continuous Integration with Travis CI Tool

In this project, we used Travis CI integrated with this repo to listen for whenever a code is merged into master to automatically trigger a process that will first run tests and of the test is successful, we then create a docker mage and then push the new image ast the latest image for the remove repo in DockerHub. DOcker helps streamline that process of Continuous Integration. For this travis project, the `.travis.yml` file is just a series of steps that define some processes that we want. We use different [phases](https://docs.travis-ci.com/user/job-lifecycle/). There is other terms in travis like jobs, builds, stage. See this [doc](https://docs.travis-ci.com/user/for-beginners/#builds-stages-jobs-and-phases)

Continuous Integration is automating the process in which code is tested, built into a docker image and deployed to a container registry. It is automating the process in which our code is packaged to something that can be deployed.

By using a CI/CD pipeline, we reduce the risk of failed deployments from user error. All of these steps (test code->build docker image->deploy to registry) can become automated. If there is an issue with the deployed code, a user can revert the changes with an older image rather than backtracking by using older code and downgrading packages and resolving dependencies.

Travis CI will build our docker image and push it to DockerHub (an image registry).

## Setting Environment variables with Travis CI

- [https://docs.travis-ci.com/user/environment-variables#defining-variables-in-repository-settings](https://docs.travis-ci.com/user/environment-variables#defining-variables-in-repository-settings)
- Also there are full list of default environment variables (see [list](https://docs.travis-ci.com/user/environment-variables#default-environment-variables)) that you can use to write your bash scripts. See eg like [here](https://docs.travis-ci.com/user/job-lifecycle/#complex-build-commands) and [here](https://stackoverflow.com/questions/37544306/travis-different-script-for-different-branch)

## Specifying Branches that we want built

For this project we used master which is default that why it was not listed. We can specify any other branch name we want. See the doc [here](https://docs.travis-ci.com/user/customizing-the-build/#safelisting-or-blocklisting-branches)

## Other Useful Articles about Travis

- [CI (Continuous Integration) with Travis CI for Golang project](https://medium.com/@classik19881/ci-continuous-integration-with-travis-ci-for-golang-project-532d1d1fc7b6)
- [Setting up Travis CI for Go project](https://wilson-tech.medium.com/setting-up-travis-ci-for-go-project-a538722f0c07)
- [Travis CI Features](https://docs.travis-ci.com/user/for-beginners/). Follow the side navigation on that page :)
