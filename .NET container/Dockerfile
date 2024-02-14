FROM mcr.microsoft.com/dotnet/sdk:8.0

WORKDIR /app

COPY *.csproj .
RUN dotnet restore

COPY . .
RUN dotnet publish -o /app

USER 10000

ENV ASPNETCORE_URLS=http://+:5000/
EXPOSE 5000

ENTRYPOINT ["dotnet", "dotnetapp.dll"]
