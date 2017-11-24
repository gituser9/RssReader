package entity;

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

    @Column(name = "TwitterScreenName")
    private String twitterScreenName;


    public long getId() {
        return id;
    }

    public String getName() {
        return name;
    }

    public String getPassword() {
        return password;
    }

    public boolean isVkNewsEnabled() {
        return vkNewsEnabled;
    }

    public String getVkLogin() {
        return vkLogin;
    }

    public String getVkPassword() {
        return vkPassword;
    }

    public String getTwitterScreenName() {
        return twitterScreenName;
    }
}
