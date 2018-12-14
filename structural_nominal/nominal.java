import java.util.Arrays;
import java.util.Comparator;

public class Main {
    private static String a;

    public static void main(String[] args) throws InterruptedException {
        Person[] array = new Person[] {
                new Person("Bob", 31),
                new Person("John", 42),
                new Person("Michael", 17),
                new Person("Jenny", 26)
        };
        Arrays.sort(array, new SortByAge());
        System.out.println(Arrays.toString(array));
    }

    public static class Person {
        public String name;
        public int age;

        public Person(String name, int age) {
            this.name = name;
            this.age = age;
        }

        @Override
        public String toString() {
            return String.format("{%s %d}", name, age);
        }
    }

    public static class SortByAge implements Comparator<Person>  {
        @Override
        public int compare(final Person p1, final Person p2) {
            return Integer.compare(p2.age, p1.age);
        }
    }
}