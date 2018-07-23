import generators.GreetingMessageGenerator;

public class Application {
  public static void main(String[] args) {
    String message = new GreetingMessageGenerator().generate();
    System.out.println(message);
  }
}
