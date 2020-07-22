# Learn Go with Tests


# Wrapping up

## Pointers & errors

### Pointers

* Go copies values when you pass them to functions/methods so if you're writing a function that needs to mutate state you'll need it to take a pointer to the thing you want to change.
* The fact that Go takes a copy of values is useful a lot of the time but sometimes you won't want your system to make a copy of something, in which case you need to pass a reference. Examples could be very large data or perhaps things you intend only to have one instance of (like database connection pools).

### nil

* Pointers ca nbe nil
* When a function returns a pointer to somethingm you need to make sure you check if it's nil or you might raise a runtime exception, the compiler won't help you here
* Usful for when you want to describe a value that could be missing

### Errors

* Errors are the way to signify failure when calling a function/method.
* By listening to our tests we concluded that check for a string in an error would result in a flaky test. So We refactored to use a meaningful value instead and this resulted in easier to test code and concluded this would be easir for users o our API too.
* This is not the end of the story with error handling, you can do more sophisticated things but this is just an intro. Later sctions will cover more strategies.
* [Don't just check errors, handle them gracefully](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)

### Create new types from existing ones

* Useful for adding more domain specific meaning to values
* Can let you implement interfaces
  
## Maps

In this section, we covered a lot. We made a full CRUD (Create, Read, Update and Delete) API for our dictionary. Throughout the process we learned how to:

* Create maps
* Search for items in maps
* Add new items to maps
* Update items in maps
* Delete items from a map
* Learned more about errors
    * How to create errors that are constants
    * Writing error wrappers


## Dependency Injection

Motivated by our tests we refactored the code so we could control where the data was written by **injecting a dependency** which allowed us to:

* **Test our code** if your can't test a function easily, it's usually because of depenencies hard-wired into a function or global state. If you have a global database connection pool for instance that is used by some kind of service layer, it is likely going to be difficult to test and they will be slow to run. DI will motivate you to inject a database dependency(via an interface) which you can then mock out with something you can control in your tests.

* **Sparate our concerns**, decoupling where the data goes from how to generate it. If you ever feel like a method/function has too many responsibilites (generating data and writing to a db? handling HTTP requests and doing domain level logic?) DI is probably going to be the tool you need.

* **Allow our code to be re-used in different contexts** The first new conext our code can be used in is inside tests. But further on if someone wants to try something new with your function they can inject their own dependencies.

## Mocking

People normally get in to a bad state when they don't listen to their tests and are not respecting the refactoring stage.

If your mocking code is becoming complicated or you are having to mock out lots of things to test something, you should listen to that bad feeling and think about your code. Usually it is a sign of

* The thing you are testing is having to do too many things (because it has too many dependencies to mock)
  * Break the module apart so it does less

* Its dependencies are too fine-grained
  * Think about how you can consolidate some of these dependencies into one meaningful module

* Your test is too concerned with implementation details
  * Favour testing expected behaviour rather than the implementation

Normally a lot of mocking points to bad abstraction in your code.

**What people see here is a weakness in TDD but it is actually a strength**, more often than not poor test code is a result of bad design or put more nicely, well-designed code is easy to test.

Ever run into this situation?

* You want to do some refactoring
* To do this you end up changing lots of tests
* You question TDD and make a post on Medium titled "Mocking considered harmful"

This is usually a sign of you testing too much implementation detail. Try to make it so your tests are testing useful behaviour unless the implementation is really important to how the system runs.

### More on TDD approach

* When faced with less trivial examples, break the problem down into "thin vertical slices". Try to get to a point where you have working software backed by tests as soon as you can, to avoid getting in rabbit holes and taking a "big bang" approach.
* Once you have some working software it should be easier to iterate with small steps until you arrive at the software you need.

### Mocking

* **Without mocking important areas of your code will be untested**. In our case we would not be able to test that our code paused between each print but there are countless other examples. Calling a service that can fail? Wanting to test your system in a particular state? It is very hard to test these scenarios without mocking.
* Without mocks you may have to set up databases and other third parties things just to test simple business rules. You're likely to have slow tests, resulting in slow feedback loops.
* By having to spin up a database or a webservice to test something you're likely to have fragile tests due to the unreliability of such services.

Once a developer learns about mocking it becomes very easy to over-test every single facet of a system in terms of the way it works rather than what it does. Always be mindful about the value of your tests and what impact they would have in future refactoring.