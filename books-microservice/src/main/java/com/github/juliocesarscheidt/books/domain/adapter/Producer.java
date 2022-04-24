package com.github.juliocesarscheidt.books.domain.adapter;

public interface Producer {

	public void sendMessage(String exchange, String replyQueueName, String messageString, Integer priority);
}
