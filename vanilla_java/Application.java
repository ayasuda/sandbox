import generators.YamlMessageGenerator;

public class Application {
  public static void main(String[] args) {
    String message = new YamlMessageGenerator().generate();
    System.out.println(message);
  }
}
