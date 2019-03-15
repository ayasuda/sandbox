package with_gradle.generators;

import org.yaml.snakeyaml.Yaml;

public class YamlMessageGenerator {
  public String generate() {
    String document = "message: Hello, YAML";
    Yaml yaml = new Yaml();
    Message message = (Message) yaml.loadAs(document, Message.class);
    return message.message;
  }

  public static class Message {
    String message;

    public String getMessage() {
      return this.message;
    }
    public void setMessage(String message) {
      this.message = message;
    }
  }
}
