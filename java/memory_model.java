public class Main {
    private static String a;

    public static void f() {
        System.out.println(a);
    }

    public static void hello() {
        a = "hello, world";
        new Thread(Main::f).start();
    }

    public static void main(String[] args) {
        hello();
    }
}