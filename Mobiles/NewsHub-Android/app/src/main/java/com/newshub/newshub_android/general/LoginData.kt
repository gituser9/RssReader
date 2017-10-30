package com.newshub.newshub_android.general


import java.io.Serializable
import com.google.gson.annotations.Expose
import com.google.gson.annotations.SerializedName

class LoginData : Serializable {

    @SerializedName("username")
    @Expose
    var username: String = ""

    @SerializedName("password")
    @Expose
    var password: String = ""

    companion object {
        private const val serialVersionUID = 5940255560625697831L
    }

}
