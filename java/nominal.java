public class Main {
    public static void main(String[] args) {
        User user = new User("foo");
        System.out.println(user);
    }

    public static class User {
        private String name;

        public User(String name) {
            this.name = name;
        }

        @Override
        public String toString() {
            return String.format("User: name = %s", this.name);
        }
    }
}
