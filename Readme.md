# GBL_ Interview Screening Code Problems_Lokanadham

## Instructions
```
git clone https://github.com/lokanadham100/restaurant
cd restaurant
go run .
```

### Major Assumptions
1. Run Time Time Complexity is important but readability, maintainability, extensibility have more priority than Algo.
2. Same WRT space complexity
3. Assuming Recommendation system is stateless, As all required data is given through args. So system explicitly wont store any user related or restaurant related data
4. Covered some edge cases of empty user behaviour etc. But for simplicity will take first cuisine/cost in case of multiple cuisine/cost end up in primary at the same time.
5. `GetRestaurantRecommendations` is thread safe. So we can safely call this concurrenlty.
6. New Restos will return only 4 restos or less(if input restos list < 4) always. And ordering is first by rating & then by date of creation. We can change this order if needed. Will need change in only one file.

### Explanation
1. Splitting the system into 9 subengines. Each one is completly isolated. This way we can add extra features or sorting etc if needed into each one individually. We can even move the subengine to seperate micro service if needed.
2. We can add one more subengine/Order into the system with new file. Just implement `SubEngine` interface.
3. Order is assuming values of 100 multiples for now. This will make the job easy for future subengines to place between existing ones just by returning some order between the range. for ex: with new `GetOrder` returning `450` new subengine can be placed before `newRestaurant` subengine.
4. Each subengine will respond with `DisplayCount`. This helps in limiting the subengine response to some number.
