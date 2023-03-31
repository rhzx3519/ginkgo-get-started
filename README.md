# Ginkgo Get Started

We use container nodes like Describe and Context to organize the different aspects of code that we are testing hierarchically. In this case we are describing our book model's ability to categorize books across two different contexts - the behavior for large books "With more than 300 pages" and small books "With fewer than 300 pages".

We use setup nodes like BeforeEach to set up the state of our specs. In this case, we are instantiating two new book models: lesMis and foxInSocks.

Finally, we use subject nodes like It to write a spec that makes assertions about the subject under test. In this case, we are ensuring that book.Category() returns the correct category enum based on the length of the book. We make these assertions using Gomega's Equal matcher and Expect syntax. You can learn much more about Gomega here - the examples used throughout these docs should be self-explanatory.
