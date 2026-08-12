# A primer on test doubles

* Test doubles is the collective noun for the different ways you can construc dependencies that you can control for a subject under test (SUT), the thing you're testing. Test doubles are often a better alternative than using the real dependency as it can avoid issues like

1. Needing the internet to use an API
2. Avoid latency and other performance issues
3. Unable to exercise non-happy path cases
4. Decoupling your build from another team's.

* Wouldn't want to prevent deployments if an engineer in another team accidentally shipped a bug

In Go, you'll typically model a dependency with an interface, then implement your version to control the behaviour in a test.
