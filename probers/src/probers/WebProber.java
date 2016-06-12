package probers;

import java.io.BufferedReader;
import java.io.DataOutputStream;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.util.logging.Level;
import java.util.logging.Logger;
import java.net.URL;

import org.apache.commons.cli.*;

/**
 * A prober, used to probe html sites and indicate if the response from the site was expected.
 * Black box monitoring.
 * run with -m=GET --targetUrl=www.google.com
 */
public class WebProber {
    private static final Logger log = Logger.getLogger(WebProber.class.getName());

    @SuppressWarnings({"deprecation", "AccessStaticViaInstance"})
    private static Options setUpCommandLineFlags() {
        // Definition Stage
        Options options = new Options();
        // TODO(pheven): instrument the prober with prometheus metrics
        // TODO(pheven): create webserver + simple page to display result of probes
        // TODO(pheven): add arg to specify frequency of probes (every 5 minutes, 20 minutes, etc...)
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
        return options;
    }

    private static void help(Options options) {
        // print out cmd line help
        HelpFormatter formatter = new HelpFormatter();
        formatter.printHelp("Main", options);
        System.exit(0);

    }

    private static String executeRequest(String requestType, String targetURL, String urlParameters){
        /**
         * executeRequest, executes the given request type against the given url with the given parameters.
         *
         * Args:
         *      requestType, the type of HTTP request to be performed (GET, POST).
         *      targetURL, the url to send the POST request.
         *      urlParameters, any parameters to send along with the url in the POST request.
         * Returns:
         *      the string response, should one exist.
         *
         */
        HttpURLConnection connection = null;
        try {
            //Create connection
            URL url = new URL(targetURL);
            connection = (HttpURLConnection)url.openConnection();
            connection.setRequestMethod(requestType);
            connection.setRequestProperty("Content-Type", "application/x-www-form-urlencoded");

            if (urlParameters != null) {
                connection.setRequestProperty("Content-Length", Integer.toString(urlParameters.getBytes().length));
            }
            else {
                connection.setRequestProperty("Content-Length", Integer.toString(0));
            }

            connection.setRequestProperty("Content-Language", "en-US");

            connection.setUseCaches(false);
            connection.setDoOutput(true);

            //Send request
            if (urlParameters != null) {
                DataOutputStream wr = new DataOutputStream (
                        connection.getOutputStream());  // http://tinyurl.com/h3o426b
                wr.writeBytes(urlParameters);
                wr.close();
            }

            //Get Response
            InputStream is = connection.getInputStream();
            BufferedReader rd = new BufferedReader(new InputStreamReader(is));
            StringBuilder response = new StringBuilder();
            String line;
            while((line = rd.readLine()) != null) {
                response.append(line);
                response.append('\n');
            }
            rd.close();
            return response.toString();
        } catch (Exception e) {
            e.printStackTrace();
            return null;
        } finally {
            if(connection != null) {
                connection.disconnect();
            }
        }
    }

    private static String executePost(String targetURL, String urlParameters) {
        String response = executeRequest("POST", targetURL, urlParameters);
        // TODO(pheven): validate the response against an expected response. this is the result of the probe.
        log.log(Level.INFO, response);
        return  response;
    }

    private static String executeGet(String targetURL, String urlParameters) {
        String response = executeRequest("GET", targetURL, urlParameters);
        // TODO(pheven): validate the response against an expected response. this is the result of the probe.
        log.log(Level.INFO, response);
        return  response;
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
            switch (method) {
                case "GET": executeGet(targetUrl, urlParameters);
                    break;
                case "POST": executePost(targetUrl, urlParameters);

                    break;
                default: executeGet(targetUrl, urlParameters);
            }
        } catch (ParseException exp) {
            System.err.println("Parsing has failed. Reason: " + exp.getMessage());
        }
    }
}
