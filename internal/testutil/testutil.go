package testutil

const SafeString = "start\t\n \rend&lt;p&gt;name&lt;/p&gt;alert(&#39;name&#39;);"
const UnsafeString = "\t\n \rstart\t\n \rend<p>name</p>alert('name');\t\n \r"
