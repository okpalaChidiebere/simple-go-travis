# Understanding Continuous Integration with Travis CI Tool

In this project, we used Travis CI integrated with this repo to listen for whenever a code is merged into master to automatically trigger a process that will first run tests and of the test is successful, we then create a docker mage and then push the new image ast the latest image for the remove repo in DockerHub. DOcker helps streamline that process of Continuous Integration. For this travis project, the `.travis.yml` file is just a series of steps that define some processes that we want. We use different [phases](https://docs.travis-ci.com/user/job-lifecycle/). There is other terms in travis like jobs, builds, stage. See this [doc](https://docs.travis-ci.com/user/for-beginners/#builds-stages-jobs-and-phases)

**Good to Know:** Our Docker Image needs and environmental variable `FAVORITE_FOOD` to be defined but when we build our image we did not bake that variable into the image like `docker build -t simple-go-travis --build-arg $FAVORITE_FOOD=${$FAVORITE_FOOD} .` for security reasons because anyone can connect into the image and access our env variable which we might not want if we want that variable to be a secret. A better outcome is to feed that env variable into the container at runtime. See this [project](https://github.com/okpalaChidiebere/mastering-docker/tree/master/debug-me#handy-docker-commands-to-debug-a-container)

Continuous Integration is automating the process in which code is tested, built into a docker image and deployed to a container registry. It is automating the process in which our code is packaged to something that can be deployed.

By using a CI/CD pipeline, we reduce the risk of failed deployments from user error. All of these steps (test code->build docker image->deploy to registry) can become automated. If there is an issue with the deployed code, a user can revert the changes with an older image rather than backtracking by using older code and downgrading packages and resolving dependencies.

Travis CI will build our docker image and push it to DockerHub (an image registry).

## Setting Environment variables with Travis CI

- [https://docs.travis-ci.com/user/environment-variables#defining-variables-in-repository-settings](https://docs.travis-ci.com/user/environment-variables#defining-variables-in-repository-settings)
- Also there are full list of default environment variables (see [list](https://docs.travis-ci.com/user/environment-variables#default-environment-variables)) that you can use to write your bash scripts. See eg like [here](https://docs.travis-ci.com/user/job-lifecycle/#complex-build-commands), [here](https://stackoverflow.com/questions/37544306/travis-different-script-for-different-branch), [here](https://stackoverflow.com/questions/58300724/travis-ci-enviroment-variables-per-branch-in-travis-yml/58391724#58391724) and [here](https://stackoverflow.com/questions/67850925/using-env-property-in-bash-script)
- One Gotcha about defining env variable is that since we want one individual build, we have to put all variable in a list like we did in this yml file. See [doc](https://docs.travis-ci.com/user/environment-variables/#defining-multiple-variables-per-item)

## Specifying Branches that we want built

For this project we used master which is default that why it was not listed. We can specify any other branch name we want. See the doc [here](https://docs.travis-ci.com/user/customizing-the-build/#safelisting-or-blocklisting-branches)

## Other Useful Articles about Travis

- [CI (Continuous Integration) with Travis CI for Golang project](https://medium.com/@classik19881/ci-continuous-integration-with-travis-ci-for-golang-project-532d1d1fc7b6)
- [Setting up Travis CI for Go project](https://wilson-tech.medium.com/setting-up-travis-ci-for-go-project-a538722f0c07)
- [Travis CI Features](https://docs.travis-ci.com/user/for-beginners/). Follow the side navigation on that page :)

## Alternative CI Tools

- [Jenkins](https://www.jenkins.io/) - most flexible but more overhead of setup
- [CircleCI](https://circleci.com/) - alternative to Travis CI with many competing features
- [AWS CodeBuild](https://aws.amazon.com/codebuild/) - integrates easily with other AWS tools
