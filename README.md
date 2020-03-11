#  Crawler in Go.  

Crawler rudimentary implementation based on Queues and go channels. 

Launcher spins of the crawler with an initial set of the URL .  
Utils consists of the Scraper and a queue data strucure where the queue URL's are persisted. 

THe URL's are fed to the Scraper which does a deep scan of the URLs to download an initial set of the URLs. This process continues till the queue size is hit. 
