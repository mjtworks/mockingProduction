package probers;

import org.apache.commons.cli.*;

import java.util.Timer;

/**
 * Run the prober at a regular time interval.
 * Command line options: -m=GET --targetUrl=www.google.com -interval=1000
 */
public class ProbeScheduler {
    @SuppressWarnings({"deprecation", "AccessStaticViaInstance"})
    private static Options setUpCommandLineFlags() {
        // Definition Stage
        Options options = new Options();
        // TODO(pheven): instrument the prober with prometheus metrics
        // TODO(pheven): create webserver + simple page to display result of probes
        options.addOption("h", "help", false, "show help.");
        options.addOption(OptionBuilder.withLongOpt( "httpMethod" )
                .withDescription( "specify the http method for the prober to use (default is GET)." )
                .hasArg()
                .withArgName("METHOD")
                .create("m"));
        options.addOption(OptionBuilder.withLongOpt( "targetUrl" )
                .withDescription( "specify the target url for the prober to send to (eg. http://www.google.com)." )
                .hasArg()
                .withArgName("URL")
                .create("t"));
        options.addOption(OptionBuilder.withLongOpt( "urlParameters" )
                .withDescription( "specify any url paramaters that should be sent in the request." )
                .hasArg()
                .withArgName("PARAMS")
                .create("p"));
        options.addOption(OptionBuilder.withLongOpt( "interval" )
                .withDescription( "specify the time interval in milliseconds for the frequency of probes. "
                        + "If not provided, a single probe is sent." )
                .hasArg()
                .create("i"));
        return options;
    }

    private static void help(Options options) {
        // print out cmd line help
        HelpFormatter formatter = new HelpFormatter();
        formatter.printHelp("Main", options);
        System.exit(0);
    }

    public static void main(String[] args) {
        Options options = setUpCommandLineFlags();
        // create parser
        CommandLineParser parser = new DefaultParser();

        try {
            // set defaults
            String method = "GET";
            String urlParameters = null;
            String targetUrl = null;
            Integer interval = null;
            // parse cmd line args
            CommandLine line = parser.parse(options, args);

            if (line.hasOption("h")) {
                help(options);
            }
            if (line.hasOption("urlParameters")) {
                urlParameters = line.getOptionValue("urlParameters");
            }
            if (line.hasOption("httpMethod")) {
                method = line.getOptionValue("httpMethod");
            } else {
                System.out.println("The httpMethod flag is required.");
                help(options);
            }
            if (line.hasOption("targetUrl")) {
                targetUrl = line.getOptionValue("targetUrl");
                if (!targetUrl.startsWith("http")) {  // best effort to put the correct protocol on the url
                    targetUrl = "http://" + targetUrl;
                }
            }
            if (line.hasOption("interval")) {
                interval = Integer.parseInt(line.getOptionValue("interval"));
            }
            WebProber prober = new WebProber(method, targetUrl, urlParameters);
            Timer time = new Timer();
            if (interval != null) {
                time.schedule(prober, 0, interval);
            } else {
                prober.run();
            }
        } catch (ParseException exp) {
            System.err.println("Parsing has failed. Reason: " + exp.getMessage());
        }

    }
}
