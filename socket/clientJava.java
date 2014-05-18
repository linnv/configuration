import java.net.*;
import java.io.*;
import java.io.DataOutputStream;

public class clientJava {
	static Socket server;

	public static void main(String[] args) throws Exception {
		server = new Socket("127.0.0.1",3333);

		DataOutputStream out = new DataOutputStream(server.getOutputStream());
		try{
		out.write("1 2 3 ");
		}catch(Exception e){
			e.printStackTrace();	
		}

		/*
		BufferedReader in = new BufferedReader(new InputStreamReader(
				server.getInputStream()));
		PrintWriter out = new PrintWriter(server.getOutputStream());
		BufferedReader wt = new BufferedReader(new InputStreamReader(System.in));

		while (true) {
			String str = wt.readLine();
			out.println(str);
			out.flush();
			if (str.equals("end")) {
				break;
			}
		}
		*/
		server.close();
	}
}
