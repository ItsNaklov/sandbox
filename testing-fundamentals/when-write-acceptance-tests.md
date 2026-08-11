# When should I write acceptance tests?

The best practice is to favour having lots of fast running unit tests and a few acceptance tests, but how do you decide when you should write an acceptance tests, vs unit tests?

It's difficult to give a concrete rule, the questions typically should be asked are:

* Is this an egde case? prefer to unit test those

* Is this something that the non-computer people talk about a lot ? Prefer to have a lot of confidence the key thing "really" works, so in the tutorial he would add an acceptance test.

* Describing a user journey, rather than a specific function? Acceptance test

* Would unit tests give me enough confidence? Sometimes taking an existing joruney that already has an acceptance test, but you're adding other functionality to deal with different scenarios due to different inputs, In this case, adding another acceptance test adds a cost but brings little value, so prefer some unit tests.
