# MagnusCareChallenge

The following folder products-api has a code for building one of the API endpoints. I worked on Endpoint#2 POST api/products/search. I will explain how my code works and the time and space complexities.


The idea behind the endpoint is to use the fields/features of a product to search for certain products. 
The process of what happens when the api endpoint is called is the following:

1. Break up the given query and store it
2. Go through the query and find the products that fit that field in the data.
3. Combine all the results that were found and apply pagination to reduce to only the needed list.
4. Return the final result.

The way the data is stored is the following different hashmaps and lists
1.ProductsMap: each product is mapped to it's productID so it is easy to find
2.BrandIndex: each brandID is mapped to a brand name
3.CategoryIndex: each categoryID is mapped to the category name

Since the data was stored as hashmaps based on the ID it became easier to find, similar to how mongoDB store's it's data. 




Time Complexity: 

With this example of a query,

       { "conditions": [
                { "type": "brandName", "values": ["Brother", "Canon"] },
                { "type": "categoryName", "values": ["Printers & Scanners"] }
          ],
          "pagination": { "from": 1, "size": 3 }
        }

1. Declaring the query and iterating through all conditions of the query will be O(N) where N is the amount of conditions. 
2. In each iteration of the conditions, the result is found by iterating through all the different values provided in the condition. In this example, ["Brother","Canon"], will have to find all Brother and Canon products.
The time complexity of that is O(M) where M is the length of values. 
3. Once all results are found, pagination needs to be applied, which I just sliced the array of all results making it O(L) where L is size of the page that was wanted in the request body. 
4. Converting the result from hashmap to a JSON will also take time, which will be O(L) whereL is the size of the result that will be returned.

So in the worst case, I believe the time complexity is O(N^2) which would only happen if there are N conditions and N values in each condition. Since, looking up in a hashmap is fast, getting the results back for output is relatively fast.

Space Complexity: 

Since each product only had a few features: brand name, category name, and title for classification, it was much easier to store the data in separate hashmaps. 
When a new product is added to the collection of products, it is added to five separate hashmaps each storing the feature. In Golang, adding values to a map is O(1 + N/k) where N is the size of the map, which for us is the number of products, and k is the resizing factor. Due to hash collisions and resizing, the hashmap isn't O(1), however it is much faster than O(N) for an array. 


The time and space complexity do depend on the language as each language has features to make things better for their purpose, so this may not be accurate in another language.


#To run the code

1. Make sure that Docker Engine is installed for your operating system. 
2. Clone this repository into your local machine.
3. Navigate into the directory that you stored it in, and run the following line
    
          docker build .
     
   This will build and install all the necessary items. Once it is complete it should create a docker image.
 4. To check for the docker image, run
        
            docker images
    
    
   which should result in something looking like this
   
                             REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
                          <none>              <none>              7f6106cea7dd        6 seconds ago       999MB
                          <none>              <none>              2f3e68699aaf        16 minutes ago      999MB
                          <none>              <none>              e4f420a5288d        20 minutes ago      999MB
                          <none>              <none>              1082a9c747e4        21 minutes ago      999MB
                          <none>              <none>              8ba00669040d        33 minutes ago      999MB
                          <none>              <none>              2f2582b12183        About an hour ago   839MB
                          
                          
  
 5. Copy that Image ID for the one that was just created.
 6. Finally run the code by calling the following
 
                docker run -p 8088:8088 <IMAGE ID>    // insert your image id here
      
    the reason it is 8088 because it will only run on your local host with that port number as requested
    
 7. Now the Gin package should have created the server and it will run on http://localhost:8088
 you can call the api endpoint with the specified query in the request body. 
 
 
 Very simple nothing more. There are ways to improve such as saving previous queries in case it is called again, or hashing the features of a product as storing an integer takes less space than a string, or more. 
