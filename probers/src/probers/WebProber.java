package probers;

import java.io.BufferedReader;
import java.io.DataOutputStream;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.URL;

import org.apache.commons.cli.*;

/**
 * A prober, used to probe html sites and indicate if the response from the site was expected.
 * Black box monitoring.
 */
public class WebProber {

    private static Options setUpCommandLineFlags() {
        // Definition Stage
        Options options = new Options();
        // TODO(pheven): see if we can do shorthand (eg -t instead of targetUrl)
        // TODO(pheven): set defaults for each of these
        // TODO(pheven): instrument the prober with prometheus metrics
        // TODO(pheven): create webserver + simple page to display result of probes
        // TODO(pheven): add arg to specify frequency of probes (every 5 minutes, 20 minutes, etc...)
        options.addOption("httpMethod", true, "specify the http method for the prober to use.");
        options.addOption("targetUrl", true, "specify the target url for the prober to send to.");
        options.addOption("urlParameters", true, "specify any url paramaters that should be sent in the request.");
        return options;
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

            connection.setRequestProperty("Content-Length",
                    Integer.toString(urlParameters.getBytes().length));
            connection.setRequestProperty("Content-Language", "en-US");

            connection.setUseCaches(false);
            connection.setDoOutput(true);

            //Send request
            DataOutputStream wr = new DataOutputStream (
                    connection.getOutputStream());
            wr.writeBytes(urlParameters);
            wr.close();

            //Get Response
            InputStream is = connection.getInputStream();
            BufferedReader rd = new BufferedReader(new InputStreamReader(is));
            StringBuilder response = new StringBuilder(); // or StringBuffer if not Java 5+
            String line;
            while((line = rd.readLine()) != null) {
                response.append(line);
                response.append('\r');
            }
            rd.close();
            // TODO(pheven): validate the response against an expected response. this is the result of the probe.
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
        return  executeRequest("POST", targetURL, urlParameters);
    }

    private static String executeGet(String targetURL, String urlParameters) {
        return  executeRequest("GET", targetURL, urlParameters);
    }

    public static void main(String[] args) {
        Options options = setUpCommandLineFlags();
        // create parser
        CommandLineParser parser = new DefaultParser();
        try {
            // parse cmd line args
            CommandLine line = parser.parse(options, args);
            String targetUrl = line.getOptionValue("targetUrl");
            String urlParameters = null;
            if (line.hasOption("urlParameters")) {
                urlParameters = line.getOptionValue("urlParameters");
            }

            if (line.hasOption("httpMethod")) {
                String method = line.getOptionValue("httpMethod");
                switch (method) {
                    case "GET": executeGet(targetUrl, urlParameters);
                        break;
                    case "POST": executePost(targetUrl, urlParameters);
                        break;
                    default: executeGet(targetUrl, urlParameters);
                }
            }
        } catch (ParseException exp) {
            System.err.println("Parsing has failed. Reason: " + exp.getMessage());
        }
        System.out.println("Hello there");
        executePost("http://www.google.com", "param1=value1&param2=value2");
    }
}
