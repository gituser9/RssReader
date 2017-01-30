package entities;

import javax.crypto.Cipher;
import javax.crypto.SecretKey;
import javax.crypto.SecretKeyFactory;
import javax.crypto.spec.IvParameterSpec;
import javax.crypto.spec.PBEKeySpec;
import javax.crypto.spec.SecretKeySpec;
import javax.persistence.*;
import java.io.Serializable;
import java.security.AlgorithmParameters;
import java.security.spec.KeySpec;


@Entity
@Table(name = "users")
public class UserEntity implements Serializable {

    private static final long serialVersionUID = -8706689714326132798L;
    private static final String salt = "DEChZn7LOdgXt6TYFAmyl3oivSqrRM";

    @Id
    @Column(name = "Id")
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private int id;

    @Column(name = "Name")
    private String name;

    @Column(name = "Password")
    private String password;

    @Column(name = "VkNewsEnabled")
    private boolean vkNewsEnabled;

    @Column(name = "VkLogin")
    private String vkLogin;


    @Column(name = "VkPassword")
    private String vkPassword;


    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public boolean getVkNewsEnabled() {
        return vkNewsEnabled;
    }

    public void setVkNewsEnabled(boolean vkNewsEnabled) {
        this.vkNewsEnabled = vkNewsEnabled;
    }

    public String getVkLogin() {
        return vkLogin;
    }

    public void setVkLogin(String vkLogin) {
        this.vkLogin = vkLogin;
    }

    public String getVkPassword() {
        try {
            /*SecretKeyFactory factory = SecretKeyFactory.getInstance("PBKDF2WithHmacSHA256");
            KeySpec spec = new PBEKeySpec(vkPassword.toCharArray(), salt.getBytes(), 65536, 256);
            SecretKey tmp = factory.generateSecret(spec);
            SecretKey secret = new SecretKeySpec(tmp.getEncoded(), "AES");
            Cipher cipher = Cipher.getInstance("AES/CBC/PKCS5Padding");
            AlgorithmParameters params = cipher.getParameters();
            byte[] iv = params.getParameterSpec(IvParameterSpec.class).getIV();
            cipher.init(Cipher.DECRYPT_MODE, secret, new IvParameterSpec(iv));

            return new String(cipher.doFinal(vkPassword.getBytes()), "UTF-8");*/

            return vkPassword;
        } catch (Exception e) {
            return null;
        }
    }

    public void setVkPassword(String vkPassword) {
        this.vkPassword = vkPassword;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }
}
