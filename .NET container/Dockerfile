FROM mcr.microsoft.com/dotnet/sdk:8.0

WORKDIR /app

COPY *.csproj .
RUN dotnet restore

COPY . .
RUN dotnet publish -o /app

USER $APP_UID

ENV ASPNETCORE_URLS=http://+:5000/
EXPOSE 5000

ENTRYPOINT ["dotnet", "dotnetapp.dll"]
