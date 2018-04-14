from ctypes import *

pfoo = cdll.LoadLibrary("./libgoo.so").add

class PyFoo(object):
    def __init__(self):
        self.foo = pfoo.New()
    
    def Bar(self):
        pfoo.Bar()


pfoo.argtypes =[]
pfoo.restype = c_char_p
print(pfoo())