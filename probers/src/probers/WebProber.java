package probers;

import java.io.BufferedReader;
import java.io.DataOutputStream;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.util.TimerTask;
import java.util.logging.Level;
import java.util.logging.Logger;
import java.net.URL;

/**
 * A prober, used to probe html sites and indicate if the response from the site was expected.
 * Black box monitoring.
 * Run from ProbeScheduler
 */
public class WebProber extends TimerTask {
    private static final Logger log = Logger.getLogger(WebProber.class.getName());
    String method;
    String targetURL;
    String urlParameters;

    public WebProber(String method, String targetURL, String urlParameters) {
        this.method = method;
        this.targetURL = targetURL;
        this.urlParameters = urlParameters;
    }
    public  static String executeRequest(String requestType, String targetURL, String urlParameters){
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

            //Send params if included
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

    @Override
    public void run() {
        // TODO(pheven): validate the response against an expected response. this is the result of the probe.
        String response = executeRequest(this.method, this.targetURL, this.urlParameters);
        log.log(Level.INFO, response);
    }
}
