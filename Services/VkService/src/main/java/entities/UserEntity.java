package entities;

import org.apache.commons.codec.binary.Base64;

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
        return vkPassword;
    }

    public String getDecryptVkPassword(String salt) {
        try {
            IvParameterSpec iv = new IvParameterSpec(salt.getBytes("UTF-8"));
            SecretKeySpec skeySpec = new SecretKeySpec(salt.getBytes("UTF-8"), "AES");

            Cipher cipher = Cipher.getInstance("AES/CBC/PKCS5PADDING");
            cipher.init(Cipher.DECRYPT_MODE, skeySpec, iv);

            byte[] original = cipher.doFinal(Base64.decodeBase64(getVkPassword()));

            return new String(original);    // decoded password
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
