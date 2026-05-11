CC = gcc
CFLAGS = -Wall -g
TARGET = c_project
SRCS = main.c math_utils.c

all: $(TARGET)
$(TARGET): $(SRCS)
	$(CC) $(CFLAGS) -o $(TARGET) $(SRCS)
	./$(TARGET)
clean:
	rm -f $(TARGET) *.o
