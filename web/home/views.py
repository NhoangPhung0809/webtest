from django.shortcuts import render
from django.http import HttpResponse, response
# Create your views here.
def index(request):
    response = HttpResponse()
    response.writelines("<h1>Xin chào mọi người </h1>")
    response.write("Đây là web nè nhe trang Home ")
    return response
