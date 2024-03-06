FROM eclipse-temurin:17-jdk-alpine
RUN addgroup --system spring && adduser --system spring --ingroup spring
COPY target/*.jar app.jar
USER spring:spring
ENTRYPOINT ["java","-jar","/app.jar"]