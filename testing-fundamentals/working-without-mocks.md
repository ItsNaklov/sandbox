In this chapter delves into the world of test doubles and explores how they influence the testing and development process. Uncover the limitations of traditional mocks, stubs, and spies and introduce a more efficient and adaptable approach using fakes and contracts.

# tl;dr

* Mocks, spies and stubs encourage you to encode assumptions of the behaviour of your dependencies ad-hocly in each test.

* These assumptions are usually not validated beyond manual checking, so they threathen your test suit's usefulness.

* Fakes and contacts give us a more sustainable method for creating test doubles with validated assumptions and better reuse than the alternatives.

* Fakes and contracts allow developers to test their systems with more realistic scenarios, improve local development experience with faster and more accurate feedback loops, and manage the complexity of evolving dependencies.
* When a project grows, though, these kinds of thes doubles can become a maintenance burden, and we should instead look to other design ideas to keep our system easy to reason and test.
