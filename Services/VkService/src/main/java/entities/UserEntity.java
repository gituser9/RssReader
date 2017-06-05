package entities;

import javax.persistence.*;
import java.io.Serializable;


@Entity
@Table(name = "users")
public class UserEntity implements Serializable {

    private static final long serialVersionUID = -8706689714326132798L;

    @Id
    @Column(name = "Id")
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private long id;

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


    public long getId() {
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

    /*public String getDecryptVkPassword(String salt) {
        String nonce = "37b8e8a308c354048d245f6d";
        String key = "AES128Key-16Char";

        if (vkPassword == null || vkPassword.isEmpty()) {
            return null;
        }

        try {
            Cipher cipher = Cipher.getInstance("AES/GCM/NoPadding");
            GCMParameterSpec spec = new GCMParameterSpec(128, nonce.getBytes());
            Key aesKey = new SecretKeySpec(key.getBytes(), "AES");
            cipher.init(Cipher.DECRYPT_MODE, aesKey, spec);
            byte[] original = cipher.doFinal(vkPassword.getBytes());

            return new String(original);    // decoded password
        } catch (Exception e) {
            return null;
        }
    }*/

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
