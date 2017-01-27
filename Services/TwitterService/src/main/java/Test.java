import twitter4j.*;

import java.util.List;


public class Test {

    public static void main(String[] args) {

        Twitter twitter = new TwitterFactory().getInstance();
        try {
            List<Status> statuses;
            String user = "google";
            statuses = twitter.getUserTimeline(user);
            /*if (args.length == 1) {
                user = args[0];
            } else {
                user = twitter.verifyCredentials().getScreenName();
                statuses = twitter.getUserTimeline();
            }*/
            System.out.println("Showing @" + user + "'s user timeline.");
            for (Status status : statuses) {
                System.out.println("@" + status.getUser().getScreenName() + " - " + status.getText());
            }
        } catch (TwitterException te) {
            te.printStackTrace();
            System.out.println("Failed to get timeline: " + te.getMessage());
            System.exit(-1);
        }


    }

}
