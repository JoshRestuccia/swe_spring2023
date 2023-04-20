### Backend

Updated backend to cryptocurrency objects and updated the documentation accordingly

Documentation: https://documenter.getpostman.com/view/25705725/2s93Y2T2Tq

### Frontend

On the front end, we decided to start working on the charts that would present data like the distribution of assets in a user's portfolio
and it's performance overall.

Along with that, front end was able to submit a form through the login that would sent the username, password, and email to the backend,
allowing for storage of user's data.

Tests were also done on the components of the front end like the separate pages that were made to see the charts, login, and homepage.
These were done by ng test and most of them failed because of the mat-toolbar and mat-card not being recognized, although the correct
declarations were made.

Unit tests done:

LoginComponent

MainpageComponent

StocksComponent

AppComponent

UserHome Component
