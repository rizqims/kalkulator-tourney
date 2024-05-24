import java.util.Scanner;

public class simpleCalc {

    public static void main (String[] args){
        //Deklarasi variabel
        double a,b, result;
        char op;
        Scanner s = new Scanner(System.in);
        //Input User
        System.out.println("===Simple Calculator===");
        System.out.print("Masukkan angka pertama: ");
        a = s.nextDouble();
        System.out.print("Masukkan operator(+, -, *, /): ");
        op = s.next().charAt(0);
        System.out.print("Masukkan angka kedua: ");
        b = s.nextDouble();

        switch(op){
            case '+':
            result = a + b;
            break;
            case '-':
            result = a - b;
            break;
            case '*':
            result = a * b;
            break;
            case '/':
                if(b!=0){
                    result = a / b;
                }else{
                    System.out.println("Tidak bisa dibagi nol");
                    return;
                }
            default:
                System.out.println("Operator salah!");
                return;
        }
        System.out.println("Hasil: "+result);
    }
}
