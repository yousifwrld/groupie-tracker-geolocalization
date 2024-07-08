## Groupie-Tracker

## Authors
1. Yousif Ahmed
2. S.Hasan
3. Isa

## Information
- Groupie Trackers consists on receiving a given API and manipulate the data contained in it, in order to create a site, displaying the information.
- This project also focuses on the creation of events/actions and on their visualization.

## Additional Features
1- Search bar:
The search bar allows users to search for specific information within the website. The program handles the following search cases:
- Artist/Band Name
- Members
- Locations
- First Album Date
- Creation Date
The search functionality is case-insensitive, and as the user types, the search bar provides typing suggestions. Each suggestion displays the individual type of the search case (e.g., "Freddie Mercury" -> "Member").

2- Filters: 
The filters allow the user to filter the artists/bands that will be shown.
incorporates the following filters:
- filter by creation date
- filter by first album date
- filter by number of members
- filter by locations of concerts

3- Geolocalization:
Mapping the different concerts locations of a certain artist/band given by the Client.
The project uses the Google Maps Goecoding API to convert addresses into geographic coordinates which are used to place markers for the concerts locations of a certain artist/band on a map.
