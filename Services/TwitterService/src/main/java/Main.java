import com.google.common.collect.Lists;
import com.twitter.hbc.ClientBuilder;
import com.twitter.hbc.core.Client;
import com.twitter.hbc.core.Constants;
import com.twitter.hbc.core.Hosts;
import com.twitter.hbc.core.HttpHosts;
import com.twitter.hbc.core.endpoint.StatusesFilterEndpoint;
import com.twitter.hbc.core.event.Event;
import com.twitter.hbc.core.processor.StringDelimitedProcessor;
import com.twitter.hbc.httpclient.auth.Authentication;
import com.twitter.hbc.httpclient.auth.BasicAuth;
import com.twitter.hbc.httpclient.auth.OAuth1;

import java.util.List;
import java.util.concurrent.BlockingQueue;
import java.util.concurrent.LinkedBlockingQueue;


public class Main {

    public static void main(String[] args) {
        /**
         * Declaring the connection information:
         */
        /** Set up your blocking queues: Be sure to size these properly based on expected TPS of your stream */
        BlockingQueue<String> msgQueue = new LinkedBlockingQueue<String>(100000);
        BlockingQueue<Event> eventQueue = new LinkedBlockingQueue<Event>(1000);

        /** Declare the host you want to connect to, the endpoint, and authentication (basic auth or oauth) */
        Hosts hosebirdHosts = new HttpHosts(Constants.STREAM_HOST);
        StatusesFilterEndpoint hosebirdEndpoint = new StatusesFilterEndpoint();

        // Optional: set up some followings and track terms
        List<Long> followings = Lists.newArrayList(1234L, 566788L);
        List<String> terms = Lists.newArrayList("twitter", "api");
        hosebirdEndpoint.followings(followings);
        hosebirdEndpoint.trackTerms(terms);


        // These secrets should be read from a config file
//    Authentication hosebirdAuth = new OAuth1("consumerKey", "consumerSecret", "token", "secret");
//        Authentication hosebirdAuth = new BasicAuth("username", "password");
        Authentication hosebirdAuth = new BasicAuth("cehazaj@rootfest.net", "x1z8YM2goPmFDhGVsCH5");


        /**
         * Creating a client:
         */
        ClientBuilder builder = new ClientBuilder()
                .name("Hosebird-Client-01")                              // optional: mainly for the logs
                .hosts(hosebirdHosts)
                .authentication(hosebirdAuth)
                .endpoint(hosebirdEndpoint)
                .processor(new StringDelimitedProcessor(msgQueue))
                .eventMessageQueue(eventQueue);                          // optional: use this if you want to process client events

        Client hosebirdClient = builder.build();

        // Attempts to establish a connection.
        hosebirdClient.connect();


        /**
         * Now, msgQueue and eventQueue will now start being filled with messages/events. Read from these queues however you like.
         */
        // on a different thread, or multiple different threads....
        try {
            while (!hosebirdClient.isDone()) {
                String msg = msgQueue.take();
                System.out.println(msg);
                // something(msg);
                // profit();
            }
        } catch (Exception e) {
            e.printStackTrace();
        }



        /**
         * You can close a connection with
         */
        hosebirdClient.stop();
    }



}
