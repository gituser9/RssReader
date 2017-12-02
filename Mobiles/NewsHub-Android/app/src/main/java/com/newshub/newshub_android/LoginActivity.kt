package com.newshub.newshub_android

import android.app.Activity
import android.content.Context
import android.content.Intent
import android.os.Bundle
import android.support.v7.app.AppCompatActivity
import android.view.View
import android.widget.Toast
import com.newshub.newshub_android.general.LoginData
import com.newshub.newshub_android.general.User
import kotlinx.android.synthetic.main.activity_login.*
import retrofit2.Call
import retrofit2.Callback
import retrofit2.Response


class LoginActivity : AppCompatActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_login)
    }


    fun login(view: View) {
        if (edtLogin.text.isEmpty() || edtPassword.text.isEmpty()) {
            Toast.makeText(this, R.string.login_data_required, Toast.LENGTH_SHORT).show()
            return
        }

        val data = LoginData()
        data.username = edtLogin.text.toString()
        data.password = edtPassword.text.toString()

        App.api?.auth(data)?.enqueue(object : Callback<User> {
            override fun onResponse(call: Call<User>, response: Response<User>) {
                val user = response.body() ?: return

                saveData(user)
            }
            override fun onFailure(call: Call<User>, t: Throwable) {
                Toast.makeText(this@LoginActivity, R.string.login_data_error, Toast.LENGTH_SHORT).show()
            }
        })
    }

    private fun saveData(user: User) {
        val preferences = getSharedPreferences("userInfo", Context.MODE_PRIVATE)
        val editor = preferences.edit()
        editor.putInt("userId", user.id)
        editor.commit()

        val intent = Intent()
        intent.putExtra("userId", user.id)
        setResult(Activity.RESULT_OK, intent)
        finish()
    }
}
