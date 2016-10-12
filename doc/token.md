<h3>imroc/imgo 添加token协议文档</h3>
添加token接口文档，用于校验用户身份

<h3>接口地址</h3>
| 接口名 | URL | 访问方式 |
| :---- | :---- | :---- |
| 添加token  | /1/admin/token/new   | POST |

<h3>返回码</h3>
| 错误码 | 描述 |
| :---- | :---- |
| 1 | 成功 |
| 65535 | 内部错误 |
| 65534 | 参数错误 |

<h3>基本返回结构</h3>
<pre>
{
    "ret": 1  //错误码
}
</pre>

<h3>例子</h3>
```sh
# uid 表示该token对应的用户id,expire表示该token的存留时长，单位为秒
curl -d "{\"uid\":10086,\"token\":\"JKF67897FDS325sdkfJK\",\"expire\":123456}" http://127.0.0.1:7172/1/admin/token/new
```
 * 返回
<pre>
{
    "ret": 1
}
</pre>

